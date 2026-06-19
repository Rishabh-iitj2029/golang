from fastapi import FastAPI,APIRouter
from typing import List
from fastapi  import Request

app = FastAPI()

router = APIRouter(prefix="/api/v1")
app.include_router(router)
# making a fake database

class Agent:
    def __init__(self,agentId:str,payment_status:str, validity:int):
        self.agentId = agentId
        self.payment_status = payment_status
        self.validity = validity

agent_database:List[Agent] = []
agent_database.append(Agent("4577","done",45))
agent_database.append(Agent("4578","done",45))


@router.get("/research/{agentId}")
async def handle_research(agentId:str, request:Request):
    isRegistered = False

    for agent in agent_database:
        if(agent.agentId == agentId):
            isRegistered = True
            break

    if not isRegistered:
        print("Agent is not registered")
        return{
            "message":"Payment required, you are not registered in my server",
            "ToPay":"700Rs"
        }
    
    payload = await request.body()
    print(payload)
    return {
        "message": "Access granted",
        "data":"This is you requested data => -------------------- "
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="127.0.0.1", port=8000, reload=True)




    
    

    
