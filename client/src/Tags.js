/*
 *  Tags.js
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
import { Link } from 'react-router-dom';

export default class Tags extends React.Component {
    render() {
        const tags = this.props.tags
            .filter(t => t.count > 0)
            .sort((a, b) => {
                if (a.name < b.name) { return -1; }
                if (a.name > b.name) { return 1; }
                return 0;
            }).map((tag) => {
                return (
                    <span>
                        <Link
                            className="badge badge-secondary"
                            to={"/tag/" + tag.slug}>
                            {tag.name}
                            {this.props.count && ` [${tag.count}]`}
                        </Link>
                    &nbsp;
                    </span>
                );
            });
        return <div>{tags}</div>;
    }
}
