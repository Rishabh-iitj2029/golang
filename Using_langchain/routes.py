from fastapi import FastAPI

app = FastAPI()
import httpx

base = "USD"
quotes = "EUR"
apiUrl = f"https://api.frankfurter.dev/v2/rates?base={base}&quotes={quotes}"
res = httpx.get(apiUrl)
res = res.json()[0]
print(f"As per the date {res["date"]}\n1 {base} = {res["rate"]} {quotes}")