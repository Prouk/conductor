package src

type Conductor struct {
	Bot    *ConductorBot
	Server *ConductorServer
	Config *Config
}

func InitConductor() *Conductor {
	c := new(Conductor)
	c.Config = new(Config)
	c.Server = new(ConductorServer)
	c.Bot = new(ConductorBot)
	return c
}
