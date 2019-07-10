//.... config.go


// local updates from macbookpro

func GetFile(){

/// get file

}




//SetEnvVarsFromConfigFiles

// get all the env vars and convert them to a map for easy referencing later.
func GetEnvVars() map[string]string {
	ret_map := map[string]string{}
	for _, pairs := range os.Environ() {
		parts := strings.SplitN(pairs, "=", 2)
		if len(parts) == 2 {
			ret_map[parts[0]] = parts[1]
		}
	}

	return ret_map
}


func SetEnvVarsFromConfigFiles(project_dir string) error {
	var err error

	// get original env vars
	orig_env_vars := GetEnvVars()

	// check the /etc/horiozn/hzn.json file that ships with the horizon-cli package
	default_config_file_dir := "/etc/horizon"
	if runtime.GOOS == "darwin" {
		default_config_file_dir = "/usr/local/etc/horizon"
	}
	configFile_pkg := filepath.Join(default_config_file_dir, DEFAULT_CONFIG_FILE)
	if configFile_pkg, err = filepath.Abs(configFile_pkg); err != nil {
		cliutils.Fatal(cliutils.CLI_GENERAL_ERROR, "Failed to get the absolute path for file %v. %v", configFile_pkg, err)
	}
	_, _, err = SetEnvVarsFromConfigFile(configFile_pkg, orig_env_vars, false)
	if err != nil {
		if !os.IsNotExist(err) {
			cliutils.Fatal(cliutils.CLI_GENERAL_ERROR, "Error set environment variable from file %v. %v", configFile_pkg, err)
		}
	} else {
		PACKAGE_CONFIG_FILE = filepath.Clean(configFile_pkg)
	}

	// check the user's configuration file  ~/.hzn/hzn.json
	configFile_user := filepath.Join(os.Getenv("HOME"), ".hzn", DEFAULT_CONFIG_FILE)
	if configFile_user, err = filepath.Abs(configFile_user); err != nil {
		cliutils.Fatal(cliutils.CLI_GENERAL_ERROR, "Failed to get the absolute path for file ~/.hzn/hzn.json. %v", err)
	}
	if configFile_user != configFile_pkg {
		_, _, err = SetEnvVarsFromConfigFile(configFile_user, orig_env_vars, false)
		if err != nil {
			if !os.IsNotExist(err) {
				cliutils.Fatal(cliutils.CLI_GENERAL_ERROR, "Error setting environment variable from file %v. %v", configFile_user, err)
			}
		} else {
			USER_CONFIG_FILE = filepath.Clean(configFile_user)
		}
	}

	// check the project's configuration file, it could be under current directory or the ./horizon directory.
	// it is usually setup by 'hzn dev service create' command
	configFile_project, err := GetProjectConfigFile(project_dir)
	if err != nil {
		cliutils.Fatal(cliutils.CLI_GENERAL_ERROR, "Error getting project level configuration file name. %v", err)
	}
	if configFile_project == "" {
		cliutils.Verbose("No project level configuration file found.")
	} else {
		if configFile_project, err = filepath.Abs(configFile_project); err != nil {
			cliutils.Fatal(cliutils.CLI_GENERAL_ERROR, "Failed to get the absolute path for file %v. %v", configFile_project, err)
		}
		if configFile_project != configFile_pkg && configFile_project != configFile_user {
			_, _, err = SetEnvVarsFromConfigFile(configFile_project, orig_env_vars, true)
			if err != nil {
				if !os.IsNotExist(err) {
					cliutils.Fatal(cliutils.CLI_GENERAL_ERROR, "Error set environment variable from file %v. %v", configFile_project, err)
				}
			} else {
				PROJECT_CONFIG_FILE = filepath.Clean(configFile_project)
			}
		}
	}

	return nil
}
