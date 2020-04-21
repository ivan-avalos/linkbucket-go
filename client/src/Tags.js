import React from 'react';
import { Link } from 'react-router-dom';

export default class Tags extends React.Component {
    render() {
        const tags = this.props.tags
            .filter(t => t.count > 0)
            .sort((a, b) => b.count - a.count).map((tag) => {
            return (
                <span>
                    <Link
                        className="badge badge-secondary"
                        to={"/tag/"+tag.slug}>
                        {tag.name}
                        {this.props.count && ` (${tag.count})`}
                    </Link>
                    &nbsp;
                </span>
            );
        });
        return <div>{tags}</div>;
    }
}