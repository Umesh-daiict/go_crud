const axios = require("axios")


const callApi = async (name,code)=>{
    const params = new URLSearchParams();
    params.append('name', name);
    params.append('code', code);
    axios.post('https://dev.29kreativ.com/recruitment/levels/',params, {
        headers: {
            'Authorization': `Bearer ${code}`,
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    })
    .then((res) => console.log(res.data))
    .catch((err) => console.error(err));
}   

callApi("umesh","f51fc18ab2e19e85a2a374ada103c2e3")
