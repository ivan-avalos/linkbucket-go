/*
 *  AppProvider.js
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
import axios from 'axios';
import React, { createContext } from 'react';
import { withRouter } from 'react-router-dom';

export const AppContext = createContext(null);

export class AppProvider extends React.Component {
    constructor(props) {
        super(props);
        this.inst = axios.create({
            baseURL: `${process.env.REACT_APP_LINKBUCKET_URL}/api`
        });
        this.config = {
            headers: {}
        };
        this.state = {
            isAuth: false,
            user: null,
            loading: false,
            errors: [],
            links: [],
            pagination: {
                total: 1,
                perPage: 20,
                currentPage: 1,
                firstPage: 1,
                lastPage: 1,
                nextPage: 1,
                prevPage: 1
            },
            tags: []
        };
        this.inst.interceptors.request.use(config => {
            this.setState({
                errors: [],
                loading: true
            });
            return config;
        }, error => {
            this.setState({ loading: false });
            this.handleError(error);
            return Promise.reject(error);
        });
        this.inst.interceptors.response.use(response => {
            this.setState({ loading: false });
            return response.data;
        }, error => {
            this.setState({ loading: false });
            this.handleError(error);
            return Promise.reject(error);
        });
    }

    componentWillMount() {
        if (!this.state.isAuth) {
            let json = localStorage.getItem('data');
            //console.log("get: " + json);
            if (json) {
                let data = JSON.parse(json);
                //console.log(data);
                if (data.isAuth) {
                    this.auth(data.user);
                }
            }
        }
    }

    componentDidUpdate() {
        let json = JSON.stringify({
            isAuth: this.state.isAuth,
            user: this.state.user
        });
        //console.log("set: " + json);
        localStorage.setItem('data', json);
    }

    auth(user) {
        this.config.headers['Authorization'] = `Bearer ${user.token}`;
        this.setState({
            isAuth: true,
            user: user
        });
    }

    unauth() {
        delete this.config.headers['Authorization'];
        this.setState({
            loaded: false,
            isAuth: false,
            user: null,
            errors: [],
            links: [],
            tags: []
        });
    }

    goHome() {
        this.props.history.replace({
            pathname: '/home',
            hash: '#',
            search: ''
        });
    }

    goBack() {
        this.props.history.goBack();
    }

    setSearch(search) {
        this.props.history.replace({
            search: search
        });
    }

    getSearch() {
        return this.props.location.search;
    }

    handleError = error => {
        if (error.response) {
            console.log(error.response);
            const data = error.response.data;
            switch (data.type) {
                case 'validation_failed':
                    this.setState({ errors: data.errors });
                    break;
                default:
                    alert(data.message);
            }
        } else if (error.request) {
            console.log(error.request);
            alert("There was an error making the request.");
        } else {
            console.log(error.config);
        }
    }

    addError(error) {
        const errors = this.state.errors;
        errors.push(error);
        this.setState({ errors: errors });
    }

    setPagination(data) {
        this.setState({
            pagination: {
                total: data["total"],
                perPage: data["per_page"],
                currentPage: data["current_page"],
                firstPage: data["first_page"],
                lastPage: data["last_page"],
                nextPage: data["next_page"],
                prevPage: data["prev_page"]
            }
        });
    }

    async register(user) {
        return this.inst.post('/register', user)
            .then(response => {
                this.auth(response.data);
                return true;
            })
            .catch(() => false);
    }

    async login(user) {
        return this.inst.post('/token', user)
            .then(response => {
                this.auth(response.data);
                return true;
            })
            .catch(() => false);
    }

    logout() {
        this.unauth();
    }

    async addLink(link) {
        this.props.history.replace({ url: '/home' });
        return this.inst.post('/link', link, this.config)
            .then(() => {
                this.getLinks();
                return true;
            })
            .catch(() => false);
    }

    async getLink(id) {
        return this.inst.get('/link/' + id, this.config)
            .then(response => {
                return response.data;
            }).catch(() => null);
    }

    async getLinks() {
        this.setState({ links: [] });
        const params = new URLSearchParams(this.getSearch());
        console.log(params.get('p'));
        return this.inst({
            method: 'get',
            url: '/link',
            params: {
                page: params.get('p') || 1,
                limit: process.env.REACT_APP_PAGINATE_LIMIT,
            },
            ...this.config
        })
            .then(response => {
                this.setPagination(response)
                this.setState({
                    links: response.data
                });
                return response.data;
            })
            .catch(() => null);
    }

    async getLinksForTag(slug) {
        this.setState({ links: [] });
        const params = new URLSearchParams(this.getSearch());
        return this.inst({
            method: 'get',
            url: '/tag/' + slug,
            params: { page: params.get('p') || 1, limit: 15 },
            ...this.config
        })
            .then(response => {
                this.setPagination(response)
                this.setState({
                    links: response.data
                });
                return response.data;
            })
            .catch(() => null);
    }

    async getLinksForSearch() {
        this.setState({ links: [] });
        const params = new URLSearchParams(this.getSearch());
        return this.inst({
            method: 'get',
            url: '/search',
            params: {
                page: params.get('p') || 1,
                limit: 15,
                query: params.get('q') || ''
            },
            ...this.config
        })
            .then(response => {
                this.setPagination(response)
                this.setState({
                    links: response.data
                });
                return response.data;
            })
            .catch(() => null)
    }

    async updateLink(id, link) {
        return this.inst.put('/link/' + id, link, this.config)
            .then(() => {
                return true;
            })
            .catch(() => false);
    }

    async deleteLink(id) {
        return this.inst.delete('/link/' + id, this.config)
            .then(() => {
                return true;
            })
            .catch(() => false);
    }

    async getTags() {
        return this.inst.get('/tag', this.config)
            .then(response => {
                this.setState({
                    tags: response.data
                });
                return true;
            })
            .catch(() => false);
    }

    async importOld(file) {
        const data = new FormData();
        data.append("links", file);
        return this.inst({
            method: 'post',
            url: '/import/old',
            data: data,
            ...this.config
        })
            .then(() => true)
            .catch(() => false);
    }

    render() {
        return (
            <AppContext.Provider
                value={{
                    state: this.state,
                    goHome: this.goHome.bind(this),
                    goBack: this.goBack.bind(this),
                    setSearch: this.setSearch.bind(this),
                    getSearch: this.getSearch.bind(this),
                    addError: this.addError.bind(this),
                    register: this.register.bind(this),
                    login: this.login.bind(this),
                    logout: this.logout.bind(this),
                    addLink: this.addLink.bind(this),
                    getLink: this.getLink.bind(this),
                    getLinks: this.getLinks.bind(this),
                    getLinksForTag: this.getLinksForTag.bind(this),
                    getLinksForSearch: this.getLinksForSearch.bind(this),
                    updateLink: this.updateLink.bind(this),
                    deleteLink: this.deleteLink.bind(this),
                    getTags: this.getTags.bind(this),
                    importOld: this.importOld.bind(this),
                }}
            >
                {this.props.children}
            </AppContext.Provider>
        );
    }
}

export const AppProviderWithRouter = withRouter(AppProvider);
