/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

import React, {Component} from 'react';
import {T, checkSession} from './Utils'

class Index extends Component {
    constructor(props) {
        super(props)
        this.state = {
            token: props.token || {},
        }
    }

    renderLogin() {
        if (!checkSession()) {
            return <a className="p-button--brand" href="/login" alt="">{T('login')}</a>
        }
    }

    render() {
        return (
            <div className="row">

                <section className="row">
                    <div className="row">
                        <div className="first">
                            <h2>{T('get-started')}</h2>
                            <ul className="p-list">
                                <li className="p-list__item is-ticked">{T('site-description1')}</li>
                                <li className="p-list__item is-ticked">{T('site-description2')}</li>
                                <li className="p-list__item is-ticked">{T('site-description3')}</li>
                                <li className="p-list__item is-ticked">{T('site-description4')}</li>
                            </ul>
                            {this.renderLogin()}
                        </div>
                    </div>
                </section>
            </div>
        );
    }
}

export default Index;