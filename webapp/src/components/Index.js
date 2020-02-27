// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

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