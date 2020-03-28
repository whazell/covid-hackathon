import React, {useState} from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  useParams
} from "react-router-dom";
import { Select, Typography, Input } from 'antd';
import {HomePage, HeaderTitle,  BusinessDetail, NewFactForm} from './components'
import './App.less';

const { Search } = Input;

const App = () => {
  const [showNewFactForm, setShowNewFactForm ] = useState(false);
  return(
  <Router>
     <section style={{ textAlign: 'center', marginTop: 48, marginBottom: 60 }}>
        <HeaderTitle/>
    <Search style={{width:400}} placeholder="Search the company" onSearch={value => console.log(value)} enterButton />  
  </section>  
    <NewFactForm/>
    <Switch>
      <Route exact path="/" children={<HomePage/>} />
      <Route path="/:id" children={<BusinessDetail />} />
    </Switch>
  </Router>)
};

export default App;
