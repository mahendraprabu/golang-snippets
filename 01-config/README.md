A sample program to read from a config/properties file during program startup to initilize program configurations.
This program uses config.json file to store the properties in JSON format.

Config struct - define your own config struct which will hold all the values read from the config file.

Function loadConfig(configFilename string) - reads the config file and returns the Config struct.
