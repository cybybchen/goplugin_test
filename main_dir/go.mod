module main_dir

go 1.22.3

require (
	px_plugin v0.0.0
)

replace px_plugin v0.0.0 => ../plugin_dir

require (
	common v0.0.0
)

replace (
	common v0.0.0 => ../common
)
