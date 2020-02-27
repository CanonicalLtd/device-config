// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.


import React, { Component } from 'react';
import {checkSession, T} from "./Utils";


const links = ['network', 'time'];

class HeaderSlim extends Component {
    constructor(props) {
        super(props)
        this.state = {};
    }

    link(l) {
        if (this.props.sectionId) {
            // This is the secondary menu
            return '/' + this.props.section + '/' + this.props.sectionId + '/' + l;
        } else {
            return '/' + l;
        }
    }

    renderLoginOut() {
        if (checkSession()) {
            return <li key={'logout'} className="p-navigation__link" role="menuitem"><a href="/logout">{T("logout")}</a></li>
        } else {
            return <li key={'login'} className="p-navigation__link" role="menuitem"><a href="/login">{T("login")}</a></li>
        }
    }

    render() {
        return (
            <header id="navigation" class="p-navigation header-slim">
                <div className="p-navigation__banner row">
                    <div className="p-navigation__logo">
                        <div className="u-vertically-center">
                            <img src="/static/images/logo.png" width="150px"  />
                        </div>
                    </div>

                    <nav className="p-navigation__nav">
                        <span className="u-off-screen"><a href="#navigation">Jump to site</a></span>
                        <ul className="p-navigation__links" role="menu">
                            {links.map((l) => {
                                var active = '';
                                if ((this.props.section === l) || (this.props.subsection === l)) {
                                    active = ' active'
                                }
                                return (
                                    <li key={l} className={'p-navigation__link' + active} role="menuitem"><a href={this.link(l)}>{T(l)}</a></li>
                                )
                            })}
                            {this.renderLoginOut()}
                        </ul>
                    </nav>
                </div>
            </header>
        );
    }
}

export default HeaderSlim;