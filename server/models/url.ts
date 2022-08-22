// import sequelize from 'sequelize';
import {
    DataTypes,
    Model,
    InferAttributes,
    InferCreationAttributes,
    CreationOptional,
    ForeignKey,
} from 'sequelize';

import User from './user';
import { sequelize } from '../index';

class Url extends Model<
    InferAttributes<Url>,
    InferCreationAttributes<Url, { omit: 'id' }>
> {
    declare id: CreationOptional<number>;
    declare originalUrl: String;
    declare newUrl: String;
    declare userId: ForeignKey<User['id']>;
}

Url.belongsTo(User);

Url.init(
    {
        id: {
            type: DataTypes.INTEGER,
            primaryKey: true,
            autoIncrement: true,
        },
        originalUrl: {
            type: DataTypes.STRING,
            validate: {
                isUrl: true,
            },
        },
        newUrl: DataTypes.STRING,
        userId: DataTypes.INTEGER,
    },
    {
        tableName: 'urls',
        sequelize,
    }
);

export default Url;
