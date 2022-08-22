import {
    Association,
    DataTypes,
    HasManyAddAssociationMixin,
    HasManyCountAssociationsMixin,
    HasManyCreateAssociationMixin,
    HasManyGetAssociationsMixin,
    HasManyHasAssociationMixin,
    HasManySetAssociationsMixin,
    HasManyAddAssociationsMixin,
    HasManyHasAssociationsMixin,
    HasManyRemoveAssociationMixin,
    HasManyRemoveAssociationsMixin,
    Model,
    ModelDefined,
    Optional,
    Sequelize,
    InferAttributes,
    InferCreationAttributes,
    CreationOptional,
    NonAttribute,
    ForeignKey,
} from 'sequelize';

import Url from './url';
import { sequelize } from '../index';

class User extends Model<
    InferAttributes<User>,
    InferCreationAttributes<User, { omit: 'id' }>
> {
    declare id: CreationOptional<number>;
    declare firstname: String;
    declare lastname: String;
    declare email: String;
    declare password: String;
    declare static associations: {
        urls: Association<User, Url>;
    };
}

User.hasMany(Url, {
    foreignKey: 'userId',
});

User.init(
    {
        id: {
            type: DataTypes.INTEGER,
            primaryKey: true,
            autoIncrement: true,
        },
        firstname: {
            type: DataTypes.STRING,
            validate: {
                isAlphanumeric: true,
                len: [2, 15],
            },
        },
        lastname: {
            type: DataTypes.STRING,
            validate: {
                isAlphanumeric: true,
                len: [2, 15],
            },
        },
        email: {
            type: DataTypes.STRING,
            unique: true,
            validate: {
                isEmail: true,
            },
        },
        password: DataTypes.STRING,
    },
    {
        tableName: 'users',
        sequelize,
    }
);


export default User;
