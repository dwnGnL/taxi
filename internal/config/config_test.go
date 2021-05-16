package config

import (
	"fmt"
	"testing"
	"time"
)

// TestInit ...
func TestInit(t *testing.T) {
	path := "../../config/config.json"
	conf, err := Init(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)

	var wantedport string = ":8085"
	var wantedTimeout time.Duration = 10000000000
	if conf.Http.Port != wantedport || conf.Http.ReadTimeout != wantedTimeout {
		t.Errorf("could open but couldn't get data \nwant port = %s timeout = %d \ngot  port = %s timeout = %d", wantedport, wantedTimeout,
			conf.Http.Port, conf.Http.ReadTimeout)
	}

	var wantedSigninKey = "key"
	var wantedExpireDuration = 5 * time.Minute

	if conf.JWT.SigningKey != wantedSigninKey || conf.JWT.ExpiredDuration != wantedExpireDuration {
		t.Errorf("could open but couldn't get data \nwant wantedSigninKey = %s wantedExpireDuration = %d \ngot  wantedSigninKey = %s wantedExpireDuration = %d",
			wantedSigninKey, wantedExpireDuration,
			conf.JWT.SigningKey, conf.JWT.ExpiredDuration)
	}
}
