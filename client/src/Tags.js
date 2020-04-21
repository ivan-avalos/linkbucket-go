import React from 'react';
import { Link } from 'react-router-dom';

export default class Tags extends React.Component {
    render() {
        const tags = this.props.tags.map((tag) => {
            return (
                <span>
                    <Link
                        className="badge badge-secondary"
                        to={"/tag/"+tag.slug}>
                        {tag.name}
                    </Link>
                    &nbsp;
                </span>
            );
        });
        return <div>{tags}</div>;
    }
}