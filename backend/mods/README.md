# Startpage modules

A module is a small golang module that is responsible for fetching, storing, and providing one or more pieces of information to be displayed on the start page.

# Interface

Each module must provide the following type and method signatures **exactly** as defined below.

*Tip:* Use the `test.sh` shell script to validate that modules are valid

## Types

**Options**

Options describes any options that you module may take in. For example, API keys.

```golang
type Options struct
```

**Instant**

Instance describes an instance of this module.

```golang
type Instance struct
```

## Methods

**Setup**

Setup should prepare this module for execution. The data directory is a dedicated directory for this module to use to store any information as needed. This directory may not exist and the module should create it if needed. The only errors returned should be fatal.

```golang
func Setup(dataDir string, options *Options) (*Instance, error)
```

**Refresh**

Refresh should refresh the data saved for this module if needed. This method is called frequently, so only pull new information if actually needed.

```golang
func (i *Instance) Refresh() error
```

**Get**

Get should return the latest cached value of the data for this module, or nil. This method should only fetch new data if the cache is expired, in which case it should use locks to only fetch one set of data.

*Note:* The return type of this method should be defined, below `T` is used as a placeholder for your modules specific type.

```golang
func (i *Instance) Get() T
```

**Teardown**

Teardown should perform any steps to tear down the module, such as closing files.

```golang
func (i *Instance) Teardown()
```
