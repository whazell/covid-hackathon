import React from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useParams
} from "react-router-dom";
import {HomePage,  BusinessDetail} from './components'
import './App.less';

const App = () => {

  return(
  <Router>

    <Switch>
      <Route exact path="/" children={<HomePage/>} />
      <Route path="/:id" children={<BusinessDetail />} />
    </Switch>
  </Router>)
};

export default App;
