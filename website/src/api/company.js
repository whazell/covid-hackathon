import axios from 'axios';

const URL = 'http://34.68.214.180/';

export const getAllCompanies = () => axios.get(`${URL}api/v1/company`);

export const getCompanyDetail = (id) => axios.get(`${URL}api/v1/company/${id}`)

export const createCompany = () => axios.post(`${URL}api/v1/company`,)

export const postCompanyFact = (id, citation, summary) => axios.post(`${URL}/api/v1/company/${id}/propose`, {
    Citaion: citation,
    Summary: summary,
    })
    

export const postFact = (citation, summary) => axios.post(`${URL}api/v1/fact/propose `, {
Citaion: citation,
Summary: summary,
})
