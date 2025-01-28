.PHONY: all clean

PLUGIN_NAME = myplugin.so
APP_NAME = myapp
PLUGIN_SOURCE = plugin/main.go
APP_SOURCE = app/main.go
OUTPUT = build

all: $(PLUGIN_NAME) $(APP_NAME)

$(PLUGIN_NAME): $(PLUGIN_SOURCE)
	go build -buildmode=plugin -o $(OUTPUT)/$(PLUGIN_NAME) $(PLUGIN_SOURCE)

$(APP_NAME): $(APP_SOURCE) $(PLUGIN_NAME)
	go build -o $(OUTPUT)/$(APP_NAME) $(APP_SOURCE)

clean:
	rm -rf $(OUTPUT)/
