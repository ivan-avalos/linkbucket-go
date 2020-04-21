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