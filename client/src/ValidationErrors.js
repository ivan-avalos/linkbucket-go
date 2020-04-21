/*
 *  ValidationErrors.js
 *  Copyright (C) 2020  Iván Ávalos <ivan.avalos.diaz@hotmail.com>
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Affero General Public License as
 *  published by the Free Software Foundation, either version 3 of the
 *  License, or (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU Affero General Public License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
import React from 'react';
import { Alert } from 'react-bootstrap';
import { AppContext } from './AppProvider';

export default class ValidationErrors extends React.Component {
    static contextType = AppContext;

    errorMessage(error) {
        switch (error.tag) {
            case 'required':
                return `Field '${error.field}' is required.`;
            case 'unique':
                return `Field '${error.field}' must be unique.`;
            case 'email':
                return `Field '${error.field}' must be a valid email.`;
            case 'url':
                return `Field '${error.field}' must be a valid URL.`;
            case 'lt':
                return `Field '${error.field}' must be less than ${error.param} characters.`;
            case 'gt':
                return `Field '${error.field}' must be more than ${error.param} characters.`;
            case 'password_match':
                return `Passwords must match.`;
            default:
                return `Unknown error in field ${error.field}`;
        }
    }

    render() {
        const errors = this.context.state.errors.map((error) => {
            return (
                <Alert variant="danger">
                    {this.errorMessage(error)}
                </Alert>
            );
        });
        return (
            <div>
                {errors}
            </div>
        );
    }
}