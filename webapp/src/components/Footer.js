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

class Footer extends Component {
    renderBullets() {
        if ((!this.props.config.custom.bullet) || (this.props.config.custom.bullet.length===0)) {
            return
        }

        let index = 0
        return (
            <ul className="p-footer__links">
                {this.props.config.custom.bullet.map(b => {
                    index++
                    return (
                        <li key={index} className="p-footer__item">
                            <a className="p-footer__link" href={b.url}>{b.text}</a>
                        </li>
                    )
                })}
            </ul>
        )
    }

    render() {
        return (
            <div id="footer">
                <div className="row footer">
                    <p>{this.props.config.custom.copyright}</p>
                    <nav className="p-footer__nav" role="navigation">
                        {
                            this.renderBullets()
                        }
                    </nav>
                </div>
            </div>
        );
    }
}

export default Footer;