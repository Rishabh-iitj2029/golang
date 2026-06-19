from langchain_groq import ChatGroq
from langchain_core.tools import tool
import os
import httpx
import requests
import config
from langchain_core.messages import HumanMessage, BaseMessage,ToolMessage

@tool(description="finds out the latest exchange rates")
def exchange_rate(base:str, quotes:str):
    apiUrl = f"https://api.frankfurter.dev/v2/rates?base={base}&quotes={quotes}"
    res = httpx.get(apiUrl)
    res = res.json()
    return f"As per the date {res[0]['date']}\n1 {base} = {res[0]['rate']} {quotes}"


@tool(description="for making reasearch")
def make_research(agentId:str,msg:str):
    apiUrl = f"http://127.0.0.1:8000/api/v1/research/{agentId}"
    body_data = {"service":msg}

    res = requests.get(apiUrl,json=body_data)
    return res.json()
    

llm = ChatGroq(model="llama-3.3-70b-versatile", temperature=0)

tools = [exchange_rate, make_research]
llm_with_tools = llm.bind_tools(tools)

def run_agent(user_prompt:str):
    msgs:list[BaseMessage] = [HumanMessage(content=user_prompt)]
    tool_mapping = {tool.name: tool for tool in tools}

    res = llm_with_tools.invoke(msgs)
    msgs.append(res)

    for tool_call in res.tool_calls:
        tool_name = tool_call["name"]
        tool_args = tool_call["args"]
        tool_id = tool_call["id"]
        print(f"Llama requested tool: '{tool_name}' with arguments: {tool_args}")

        target_tool = tool_mapping[tool_name]
        tool_output = target_tool.invoke(tool_args)

            
        print(f"Tool returned: {tool_output}")

        msgs.append(
                ToolMessage(
                    content=str(tool_output), 
                    tool_call_id=tool_id
                )
            )

run_agent("what is current exchange rate between USD and EUR")
print("\n-----------------------------------------------------------------------------------------------------------------------------------")
run_agent("research about the blockchain network, my agent id is 4578")