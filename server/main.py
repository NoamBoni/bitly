from typing import Union
import os
from dotenv import load_dotenv
from fastapi import FastAPI
from sqlalchemy import create_engine, sessionmaker


load_dotenv()
app = FastAPI()

connection_string = os.getenv("DB_URI").replace("postgres", "postgresql")
engine = create_engine(connection_string , echo=True)
Session = sessionmaker(bind=engine)
session = Session()



@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}