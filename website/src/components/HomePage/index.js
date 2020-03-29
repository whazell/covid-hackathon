import React, {useState} from 'react';
import { Input, Button} from 'antd';
import {BusinessCard, HeaderTitle, NewFactForm} from '../index'
import { useEffectAsync } from '../../util';
import { getAllCompanies } from '../../api/company';
import '../../App.less'

import { motion } from "framer-motion";

const { Search } = Input;

const postVariants = {
  initial: { scale: 0.96, y: 30, opacity: 0 },
  enter: { scale: 1, y: 0, opacity: 1, transition: { duration: 0.5, ease: [0.48, 0.15, 0.25, 0.96] } },
  exit: {
    scale: 0.6,
    y: 100,
    opacity: 0,
    transition: { duration: 0.2, ease: [0.48, 0.15, 0.25, 0.96] }
  }
};


const HomePage = () => {
  const [ companies, setCompanies] = useState([])
  useEffectAsync(async () => {
    const res = await getAllCompanies();
    setCompanies(res.data);
  }, [])
  const [showNewFactForm, setShowNewFactForm ] = useState(false);
  return(
    <div>
           <section style={{ textAlign: 'center', marginTop: 48, marginBottom: 60 }}>
        <HeaderTitle/>
    <Search style={{width:400}} placeholder="Search the company" onSearch={value => console.log(value)} enterButton />  
    <Button style={{paddingLeft:'20px', marginLeft:'20px'}} onClick={() => {
      setShowNewFactForm(true)
    }}>Submit New Fact</Button>
  </section>  
    <NewFactForm visible={showNewFactForm} onCancel={() => setShowNewFactForm(false)} />
    <motion.div
    initial="initial"
    animate="enter"
    exit="exit"
    variants={{ exit: { transition: { staggerChildren: 0.1 } } }}
  >
    <div className="company-grid">
    {companies.map(company =>
      <motion.div variants={postVariants}>
           <motion.div whileHover="hover" variants={{ hover: { scale: 0.96 } }}>
           <BusinessCard company={company}/>
                  </motion.div>
      </motion.div>
      )}
    </div> 
    </motion.div>
    </div>
   )
};

export default HomePage;
