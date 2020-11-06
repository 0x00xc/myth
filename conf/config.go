/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/28 13:51
 */
package conf

type Config struct {
	AppPort        int      `json:"app_port"        toml:"app_port"`        //
	AppCert        []string `json:"app_cert"        toml:"app_cert"`        //
	DatabaseName   string   `json:"database_name"   toml:"database_name"`   //
	DatabaseSource string   `json:"database_source" toml:"database_source"` //
	SMTPHost       string   `json:"smtp_host"       toml:"smtp_host"`
	SMTPPort       int      `json:"smtp_port"       toml:"smtp_port"`
	SMTPUser       string   `json:"smtp_user"       toml:"smtp_user"`
	SMTPPass       string   `json:"smtp_pass"       toml:"smtp_pass"`
}

func (c *Config) Secure() bool {
	return len(c.AppCert) >= 2 && c.AppCert[0] != "" && c.AppCert[1] != ""
}
