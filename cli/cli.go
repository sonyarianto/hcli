package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"regexp"
	"sync"
	"sort"

	"github.com/olekukonko/tablewriter"
)

type CliProgram struct {
	name string `json:"name"`
	command string `json:"command"`
	version string `json:"version"`
}

func DetectCli(name string, command string, regexPatterns interface{}, cliPrograms *[]CliProgram) {
	args := strings.Split(command, " ")

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	out, err := cmd.Output()
	if err != nil {
		return
	}

	match := ""

	switch rp := regexPatterns.(type) {
	case []string:
	for _, regexPattern := range rp {
		re := regexp.MustCompile(regexPattern)
		match = re.FindString(string(out))
		match = strings.TrimSpace(match)

		if match != "" {
			break
		}
	}
	case string:
		re := regexp.MustCompile(rp)
		match = re.FindString(string(out))
		match = strings.TrimSpace(match)
	}

	*cliPrograms = append(*cliPrograms, CliProgram{name, command, match})
}

func DetectAvailableCli() {
	var wg sync.WaitGroup
	var cliPrograms []CliProgram

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Go", "go version", []string{`(\d+\.\d+\.\d+)`, `(\d+\.\d+)`}, &cliPrograms)
	}()
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("PHP", "php -v", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Node", "node -v", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("npm", "npm -v", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Docker Client", "docker -v", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Git", "git --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Composer", "composer --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Deno", "deno --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("rustup", "rustup --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("cargo", "cargo --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("ffmpeg", "ffmpeg -version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("ffprobe", "ffprobe -version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Python", "python3 --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Python", "python --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("pip", "pip3 --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("pip", "pip --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Ruby", "ruby --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("RubyGems", "gem --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Java", "java --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Gradle", "gradle --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Maven", "mvn --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("MySQL Client", "mysql --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("PostgreSQL Client", "psql --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("SQLite3 Client", "sqlite3 --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("MongoDB Client", "mongo --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Redis Client", "redis-cli --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("RabbitMQ Client", "rabbitmqctl --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Erlang", "erl -eval 'erlang:display(erlang:system_info(otp_release)), halt().'  -noshell", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Elixir", "elixir --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()
		
	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Nginx", "nginx -v", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Apache", "httpd -v", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("GitHub CLI", "gh --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("npx", "npx --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Yarn", "yarn --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("pnpm", "pnpm --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("PowerShell Core", "pwsh --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Windows PowerShell", "'{0}.{1}.{2}' -f $PSVersionTable.PSVersion.Major, $PSVersionTable.PSVersion.Minor, $PSVersionTable.PSVersion.Build", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("nvm", "nvm --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Homebrew", "brew --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Windows Package Manager", "winget --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Chocolatey", "choco --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		DetectCli("Visual Studio Code", "code --version", `(\d+\.\d+\.\d+)`, &cliPrograms)
	}()
	
	wg.Wait()

	sort.SliceStable(cliPrograms, func(i, j int) bool {
		return strings.ToLower(cliPrograms[i].name) < strings.ToLower(cliPrograms[j].name)
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Version", "Command"})
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_LEFT})

	for _, cliProgram := range cliPrograms {
		table.Append([]string{cliProgram.name, cliProgram.version, cliProgram.command})		
	}

	table.Render()

	fmt.Println();
}