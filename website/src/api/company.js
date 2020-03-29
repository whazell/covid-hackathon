import axios from 'axios';

const URL = 'http://34.68.214.180/';

export const getAllCompanies = () => axios.get(`${URL}api/v1/company`);

export const postCompanyFact = (id, citation, summary) => axios.post(`${URL}api/v1/compnay/${id}`, {
    Citaion: citation,
    Summary: summary,
    })
    

export const postFact = (citation, summary) => axios.post(`${URL}api/v1/fact`, {
Citaion: citation,
Summary: summary,
})
