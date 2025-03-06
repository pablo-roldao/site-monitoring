# Site Monitoring

A simple website monitoring project developed in Go. This program allows you to monitor a list of websites and log their HTTP response statuses in a log file.

## Features

- **Website Monitoring:** Periodically checks the websites listed in the `sites.txt` file and displays the HTTP status of each.
- **Log Recording:** Saves the information from each monitoring cycle (date, URL, and status) in the `log.txt` file.
- **Interactive Menu:** Allows you to start monitoring, view the logs, or exit the program in a simple and intuitive way.

## Requirements

- [Go](https://golang.org/doc/install) installed on your machine (version 1.16 or later).

## Project Structure

- `main.go`: The main source code file of the project.
- `sites.txt`: A text file containing the list of URLs to be monitored (one site per line).
- `log.txt`: A file where the monitoring results are logged.

## How to Use

1. **Clone the repository:**

   ```bash
   git clone https://github.com/pablo-roldao/site-monitoring.git
   cd site-monitoring
   ```

2. **Configure the `sites.txt` file:**

   Add the URLs you want to monitor, one per line. Example:

   ```
   https://archlinux.org
   https://www.github.com
   ```

3. **Run the program:**
   Run directly
   ```bash
   go run main.go
   ```
   or compile and run
   ```bash
   go build main.go
   ./main
   ```

4. **Interact with the menu:**

   - **1:** Starts the website monitoring.
   - **2:** Displays the monitoring logs.
   - **0:** Exits the program.

## Customization

You can adjust the monitoring parameters by changing the constants defined at the beginning of the `main.go` file:

- `monitoringQuantity`: Defines the number of monitoring cycles.
- `monitoringDelay`: Defines the delay (in seconds) between each monitoring cycle.
- `sitesPath`: Defines the path to the sites file.
- `logsPath`: Defines the path to the logs file.

## License

This project is licensed under the [MIT License](LICENSE). Please refer to the `LICENSE` file for more details.
