# GoldenBoy

## How to run

* Install Golang: https://go.dev/doc/install

* Install dependencies:
    ```
    cd GoldenBoy
    go get
    ```

* Download and unzip chromedriver and chrome-linux64:
    ```
    chrome-linux64 visit: https://edgedl.me.gvt1.com/edgedl/chrome/chrome-for-testing/119.0.6045.105/linux64/chrome-linux64.zip

    chromedriver visit: https://edgedl.me.gvt1.com/edgedl/chrome/chrome-for-testing/119.0.6045.105/linux64/chromedriver-linux64.zip
    ```

* Compile and run KpScraper:
    ```
    go build cmd/KpScraper/kpscraper.go && ./kpscraper
    ```
* Compile and run moneyMaker:
    ```
    go build cmd/moneyMaker/moneyMaker.go && ./moneyMaker
    ```