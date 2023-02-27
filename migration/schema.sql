DROP TABLE  IF EXISTs Balance 


CREATE TABLE Balance(
    User_ID SERIAL PRIMARY KEY,
    Balance VARCHAR(244),
    Currency VARCHAR(244),
)