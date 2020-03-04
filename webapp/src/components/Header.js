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
import HeaderSlim from "./HeaderSlim";
import {T} from './Utils'


class Header extends Component {
    render() {
        return (
            <div>
                <HeaderSlim section={this.props.section} subsection={this.props.subsection} sectionId={this.props.sectionId} />
                <section className="p-strip--image is-dark header">
                    <div className="row">
                        <div className="col-5 title">
                            <h1>{T('title')}</h1>
                            <p>{T('subtitle')}</p>
                        </div>
                    </div>
                </section>
            </div>
        );
    }
}

export default Header;
