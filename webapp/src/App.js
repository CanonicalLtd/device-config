// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

import React, { Component } from 'react';
import Header from './components/Header';
import HeaderSlim from './components/HeaderSlim';
import Index from './components/Index';
import Footer from './components/Footer';
import Login from './components/Login';
import {parseRoute, formatError} from './components/Utils'
import api from './components/api'

import createHistory from 'history/createBrowserHistory'
import Network from "./components/Network";
import Proxy from "./components/Proxy";
const history = createHistory()

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
        location: history.location,
        token: props.token || {},
        proxy: {},
    }
  }

  handleNavigation(location) {
    this.setState({ location: location })
    window.scrollTo(0, 0)
  }

  // Get the data that's conditional on the route
  updateDataForRoute() {
      const r = parseRoute()

      if(r.section==='network') {this.getNetworkConfig()}
  }

  renderLogin() {
      return <Login />
  }

  renderNetwork(sectionId, subsection) {
      if (!sectionId) {
          return <Network />
      }
  }

  renderProxy(sectionId, subsection) {
    if (!sectionId) {
        return <Proxy />
    }
  }

  render() {
    const r = parseRoute()
    console.log(r)

    return (
        <div className="App">
          {r.section===''? <Header section={r.section} subsection={r.subsection} sectionId={r.sectionId} /> : ''}
          {r.section!==''? <HeaderSlim section={r.section} subsection={r.subsection} sectionId={r.sectionId} /> : ''}

          <div className="content row">
            {r.section===''? <Index /> : ''}
            {r.section==='login'? this.renderLogin() : ''}
            {r.section==='network'? this.renderNetwork(r.sectionId, r.subsection) : ''}
            {r.section==='proxy'? this.renderProxy(r.sectionId, r.subsection) : ''}
          </div>

          <Footer />
        </div>
    );
  }

}

export default App;
