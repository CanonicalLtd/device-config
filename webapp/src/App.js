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
import Index from './components/PageIndex';
import Footer from './components/Footer';
import Login from './components/PageLogin';
import {getLanguage, parseRoute, saveLanguage} from './components/Utils'

import Network from "./components/PageNetwork";
import Proxy from "./components/PageProxy";
import Time from "./components/PageTime";
import Snaps from "./components/PageSnaps";
import Services from "./components/PageServices";
import Settings from "./components/PageSettings";

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
        language: getLanguage(),
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
            {r.section==='settings'? <Settings config={this.props.config} /> : ''}
          </div>

          <Footer />
        </div>
    );
  }

}

export default App;
