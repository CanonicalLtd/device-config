// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.


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
