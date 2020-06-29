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
import {Select} from "@canonical/react-components";


let links = ['network','time', 'settings'];

class HeaderSlim extends Component {
    constructor(props) {
        super(props)
        this.state = {};

        if (this.props.config.snapControl) {
            links.unshift('services','snaps','proxy')
        }
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

    renderLanguage() {
        let language = getLanguage()
        let languages = Constants.languages.map(lang => {
            return {value: lang, label: T(lang)}
        })

        return (
            <li className="p-navigation__link">
                <form id="lang-form">
                    <Select onChange={this.handleLanguageChange} options={languages} defaultValue={language} />
                </form>
            </li>
        )
    }

    render() {
        return (
            <header id="navigation" className="p-navigation header-slim">
                <div className="p-navigation__banner row">
                    <div className="p-navigation__logo">
                        <div className="u-vertically-center">
                            <a href="/" className="p-navigation__link">
                                <img src="/static/custom/logo.png" alt="ubuntu" />
                            </a>
                        </div>
                        <div className="system">
                            <span>{this.props.config.snapVersion}</span>
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
                            {this.renderLanguage()}
                        </ul>
                    </nav>
                </div>
            </header>
        );
    }
}

export default HeaderSlim;