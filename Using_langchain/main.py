from fastapi import FastAPI,APIRouter, status
from typing import List
from fastapi  import Request
import datetime

DATE_FORMAT = "%Y-%m-%d %H:%M:%S"

app = FastAPI()

router = APIRouter(prefix="/api/v1")
app.include_router(router)
# making a fake database

# class Agent:
#     def __init__(self,agentId:str,payment_status:str, validity:int):
#         self.agentId = agentId
#         self.payment_status = payment_status
#         self.validity = validity

# agent_database:List[Agent] = []
# agent_database.append(Agent("4575","done",45))
# agent_database.append(Agent("4578","done",45))


@router.get("/research/")
async def handle_research(request:Request):
    # isRegistered = False

    # for agent in agent_database:
    #     if(agent.agentId == agentId):
    #         isRegistered = True
    #         break
    
    payload = await request.body()
    # if not isRegistered:
    #     print("Agent is not registered")
    #     return{
    #         "message":"Payment required, you are not registered in my server",
    #         "ToPay":"700Rs",
    #         "status": status.HTTP_402_PAYMENT_REQUIRED
    #     }
    
    with open("paymentWallet.txt", "r") as f:
        lines = f.readlines()
        payTime = datetime.datetime.strptime(lines[1], DATE_FORMAT)
        current = datetime.datetime.now()
        print(payload)
        if((current-payTime).total_seconds() <= 30):
            return {
             "message": "Access granted",
             "data":"This is you requested data => -------------------- ",
             "status": status.HTTP_200_OK
             }
        else:
            return{
            "message":"Payment required, your last payment is expired",
            "ToPay":"25",
            "status": status.HTTP_402_PAYMENT_REQUIRED
        }


    # print(payload)
    # return {
    #     "message": "Access granted",
    #     "data":"This is you requested data => -------------------- "
    # }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="127.0.0.1", port=8000, reload=True)




    
    

    
