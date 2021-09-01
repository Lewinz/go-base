package cron

import (
	"log"

	"github.com/Lewinz/golang_utils/logger/phuslog"
	"github.com/robfig/cron/v3"
)

// CronJob cron job instance
var CronJob *Cron

// Cron type
type Cron struct {
	croner *cron.Cron
	config *Config
	logger *phuslog.Logger
}

// Config ..
type Config struct {
	Enable bool                `mapstructure:"enable"`
	Specs  map[string][]string `mapstructure:"specs"` // one spec to multi events
}

// Events 返回定时任务配置
func (config *Config) Events() (res map[string]string) {
	res = make(map[string]string)
	for spec, events := range config.Specs {
		for _, event := range events {
			res[event] = spec
		}
	}
	return
}

// SetUp set instance
func SetUp(job *Cron) {
	CronJob = job
}

// Event func
type Event func()

// Start cron
func (c *Cron) Start() {
	c.croner.Start()
}

// Stop cron
func (c *Cron) Stop() {
	c.croner.Stop()
}

// NewCron server init
func NewCron(config *Config, logger log.Logger) (c *Cron) {
	c = &Cron{
		// cron V3 默认不支持秒单位，cron.WithSeconds 设置为秒单位
		croner: cron.New(cron.WithSeconds()),
		config: config,
		logger: phuslog.CornLog,
	}

	// config enable cron or need manual operation start cron
	if config != nil && config.Enable {
		c.Start()
	}

	return
}

// AddEvent name is the unique key and event must define in cron config
func (c *Cron) AddEvent(name string, event Event) {
	// config enable cron or need manual operation start cron
	if c.config != nil && c.config.Enable {

		if spec, ok := c.config.Events()[name]; ok {
			// fmt.Info("add event ", name, spec)
			c.addEvent(name, spec, event)
		}
	}
}

// addEvent name is the unique key and event must define in cron config
func (c *Cron) addEvent(name string, spec string, event Event) {
	c.croner.AddFunc(spec, func() {
		defer func() {
			if p := recover(); p != nil {
				switch i := p.(type) {
				case error:
					c.logger.Errorf("panic in crontab func name %v,spec %v,error %v\n", name, spec, i.Error())
				default:
					c.logger.Errorf("panic in crontab func name %v,spec %v,error %v\n", name, spec, i)
				}
			}
		}()
		c.eventRunner(name, spec, event)
	})
}

// eventRunner make cron a distributed run task
func (c *Cron) eventRunner(name string, spec string, event Event) {
	// start := time.Now()
	// c.logger.Infof("time: %v ,start do event(%v)", start, name)
	event()
	// c.logger.Infof("spend: %v ,end do event(%v)", time.Since(start), name)
}
