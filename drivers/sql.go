package drivers

import (
	"database/sql"

	"github.com/flashmob/go-guerrilla/backends"
	"github.com/flashmob/go-guerrilla/mail"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type sqlLiteConfig struct {
	Path string `json:"sql_lite_path"`
}

func newDBStore(config *sqlLiteConfig) error {

	_, err := sql.Open("sqlite3", config.Path)
	if err != nil {
		return err
	}

	// statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS mails (mail_id INTEGER PRIMARY KEY")
	// if err != nil {
	// 	return err
	// }

	//statement.Exec()

	return nil
}

var SqlLiteProcessor = func() backends.Decorator {

	config := &sqlLiteConfig{}
	// our initFunc will load the config.
	initFunc := backends.InitializeWith(func(backendConfig backends.BackendConfig) error {
		configType := backends.BaseConfig(&sqlLiteConfig{})
		bcfg, err := backends.Svc.ExtractConfig(backendConfig, configType)
		if err != nil {
			return err
		}
		config = bcfg.(*sqlLiteConfig)

		//create account
		err = newDBStore(config)
		if err != nil {
			return err
		}

		return nil
	})
	// register our initializer
	backends.Svc.AddInitializer(initFunc)

	return func(p backends.Processor) backends.Processor {
		return backends.ProcessWith(
			func(e *mail.Envelope, task backends.SelectTask) (backends.Result, error) {
				if task == backends.TaskValidateRcpt {

					// if you want your processor to validate recipents,
					// validate recipient by checking
					// the last item added to `e.RcptTo` slice
					// if error, then return something like this:
					/* return backends.NewResult(
					   response.Canned.FailNoSenderDataCmd),
					   backends.NoSuchUser
					*/
					// if no error:
					log.Println("Not Validating Anything")
					return p.Process(e, task)
				} else if task == backends.TaskSaveMail {

					log.Debug("Storing Data To File")
					//store to db

					log.Debug("From : ", e.MailFrom.String())
					log.Debug("Subject : ", e.Subject)

					return p.Process(e, task)
				}
				return p.Process(e, task)
			},
		)
	}
}
