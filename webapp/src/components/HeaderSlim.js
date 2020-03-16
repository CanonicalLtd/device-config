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

import React, { Component } from 'react';
import {checkSession, getLanguage, T} from "./Utils";
import Constants from "./constants";


const links = ['services','network','time'];

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

    handleLanguageChange = (e) => {
        this.props.changeLanguage(e.target.value)
    }

    renderLoginOut() {
        if (checkSession()) {
            return <li key={'logout'} className="p-navigation__link" role="menuitem"><a href="/logout">{T("logout")}</a></li>
        } else {
            return <li key={'login'} className="p-navigation__link" role="menuitem"><a href="/login">{T("login")}</a></li>
        }
    }

    renderProxy() {
        if ((!this.props.config) || (!this.props.config.manageProxy)) {
            return
        }

        let l = 'proxy'
        let active = '';
        if ((this.props.section === l) || (this.props.subsection === l)) {
            active = ' active'
        }
        return (
            <li key={l} className={'p-navigation__link' + active} role="menuitem"><a href={this.link(l)}>{T(l)}</a></li>
        )
    }

    renderLanguage() {
        let language = getLanguage()
        let languages = Constants.languages
        return (
            <li className="p-navigation__link">
                <form id="lang-form">
                    <select onChange={this.handleLanguageChange}>
                    {languages.map(lang => {
                        let selected = lang===language ? 'selected' : ''
                        return (
                            <option key={lang} value={lang} selected={selected}>{T(lang)}</option>
                        )
                    })}
                    </select>
                </form>
            </li>
        )
    }

    render() {
        return (
            <header id="navigation" class="p-navigation header-slim">
                <div className="p-navigation__banner row">
                    <div className="p-navigation__logo">
                        <div className="u-vertically-center">
                            <a href="/" className="p-navigation__link">
                                <img src="/static/images/logo.png" alt="ubuntu" />
                            </a>
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
                            {this.renderProxy()}
                            {this.renderLoginOut()}
                            {this.renderLanguage()}
                        </ul>
                    </nav>
                </div>
            </header>
        );
    }
}

export default HeaderSlim;