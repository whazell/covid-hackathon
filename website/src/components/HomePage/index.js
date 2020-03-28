import React, {useState} from 'react';
import {BusinessCard} from '../index'
import { useEffectAsync } from '../../util';
import { getAllCompanies } from '../../api/company';
import '../../App.less'


const HomePage = () => {
  const [ companies, setCompanies] = useState([])
  useEffectAsync(async () => {
    const res = await getAllCompanies();
    setCompanies(res.data);
  }, [])
  return(
  <div>
    <div className="company-grid">
    {companies.map(company => <BusinessCard company={company}/>)}
    </div> 
  </div>
   )
};

export default HomePage;
