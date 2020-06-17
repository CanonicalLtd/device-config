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
import Header from './components/Header';
import HeaderSlim from './components/HeaderSlim';
import Index from './components/Index';
import Footer from './components/Footer';
import Login from './components/Login';
import {getLanguage, parseRoute, saveLanguage} from './components/Utils'

import Network from "./components/Network";
import Proxy from "./components/Proxy";
import Time from "./components/Time";
import Snaps from "./components/Snaps";
import Services from "./components/Services";

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
        language: getLanguage(),
        token: props.token || {},
        proxy: {},
    }
  }

  changeLanguage = (l) => {
      saveLanguage(l)
      this.setState({language: l})
  }

  renderLogin() {
      return <Login />
  }

  renderNetwork(sectionId, subsection) {
      if (!sectionId) {
          return <Network config={this.props.config} />
      }
  }

  renderProxy(sectionId, subsection) {
    if (!sectionId) {
        return <Proxy />
    }
  }

  renderTime(sectionId, subsection) {
    if (!sectionId) {
        return <Time />
    }
  }

    renderSnaps(sectionId, subsection) {
        if (!sectionId) {
            return <Snaps />
        }
    }

    renderServices(sectionId, subsection) {
        if (!sectionId) {
            return <Services />
        }
    }

  render() {
    const r = parseRoute()

    return (
        <div className="App">
          {r.section===''? <Header section={r.section} subsection={r.subsection} sectionId={r.sectionId} config={this.props.config} changeLanguage={this.changeLanguage} /> : ''}
          {r.section!==''? <HeaderSlim section={r.section} subsection={r.subsection} sectionId={r.sectionId} config={this.props.config} changeLanguage={this.changeLanguage} /> : ''}

          <div className="content row">
            {r.section===''? <Index /> : ''}
            {r.section==='login'? this.renderLogin() : ''}
            {r.section==='network'? this.renderNetwork(r.sectionId, r.subsection) : ''}
            {r.section==='proxy'? this.renderProxy(r.sectionId, r.subsection) : ''}
            {r.section==='time'? this.renderTime(r.sectionId, r.subsection) : ''}
            {r.section==='services'? this.renderServices(r.sectionId, r.subsection) : ''}
            {r.section==='snaps'? this.renderSnaps(r.sectionId, r.subsection) : ''}
          </div>

          <Footer />
        </div>
    );
  }

}

export default App;
