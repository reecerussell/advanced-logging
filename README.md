# Advanced Logging

<img src="https://media.giphy.com/media/W2x2NyL5Of6JW/giphy.gif" align="right" width="260" />

Advanced Logging is an expandable, plugin based logging library, written in Go.

## Configuration

The library is configured through a JSON file, called `logging.json`.

Plugins are defined and configured using the following properties:

-   `printPrefix` - a string that will prefix all print calls.
-   `panicPrefix` - a string that will prefix all panic calls.
-   `fatalPrefix` - a string that will prefix all fatal calls.
-   `plugins` - an array of plugin definitions.

```json
{
	"printPrefix": "INFO",
	"panicPrefix": "CRASH",
	"fatalPrefix": "ERROR",
	"plugins": []
}
```

### Plugins

Each plugin must be defined in the config file. A plugin can contain the following properties:

-   `name` - a required field, used to identify each plugin.
-   `enabled` - a required boolean value, determining if it's enabled.
-   `config` - an object with plugin specific, configuration.

```json
{
    "name": "log",
    "enabled": true,
    "config": {}
},
```

#### Built-in Plugins

Advanced Logging comes with two built-in plugins. One for writing to files, and the other for logging to the console.

Each of the plugins have their own specific configuration properties, for example:

```json
[
	{
		"name": "log",
		"enabled": true,
		"config": {}
	},
	{
		"name": "file",
		"enabled": true,
		"config": {
			"outputDir": "logs"
		}
	}
]
```

As seen above, the log plugin doesn't require any configuration, whereas the file plugin requires an output directory.
