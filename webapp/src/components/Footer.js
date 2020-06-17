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
import {T} from './Utils'

class Footer extends Component {
    render() {
        return (
            <div id="footer">
                <div className="row footer">
                    <p>{T('copyright')}</p>
                    <nav className="p-footer__nav" role="navigation">
                        <ul className="p-footer__links">
                            <li className="p-footer__item">
                                <a className="p-footer__link" href="https://ubuntu.com/legal">{T('legal')}</a>
                            </li>
                            <li className="p-footer__item">
                                <a className="p-footer__link" href="https://ubuntu.com/legal/data-privacy">{T('privacy')}</a>
                            </li>
                            <li className="p-footer__item">
                                <a className="p-footer__link" href="https://github.com/CanonicalLtd/device-config/issues/new">{T('report-a-bug')}</a>
                            </li>
                        </ul>
                    </nav>
                </div>
            </div>
        );
    }
}

export default Footer;