import express from 'express';
import { Sequelize } from 'sequelize';
import dotenv from 'dotenv';

const app = express();

app.use(express.json());

dotenv.config({ path: './.env' });

const sequelize = new Sequelize(process.env.DB_URI);

sequelize
    .authenticate()
    .then(_ => console.log('Connection has been established successfully.'))
    .catch(err => console.log(err));

export { sequelize };

const port = 8000;
app.listen(port, () => {
    console.log(`application is running on port ${port}.`);
});
