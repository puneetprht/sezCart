import cookie from 'js-cookie';
import {parseCookies} from 'nookies';
import {useRouter} from 'next/router';
import axios from '../src/service/axios';
import {useState} from 'react';
import 'react-toastify/dist/ReactToastify.css';
import { ToastContainer, toast } from 'react-toastify';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons';

import styles from '../styles/Index.module.css'

export default function Home(props) {
  const router  = useRouter()

  const [nameSignUp, setNameSignUp] = useState('');
  const [emailSignUp, setEmailSignUp] = useState('');
  const [passwordSignUp, setPasswordSignUp] = useState('');
  const [emailLogin, setEmailLogin] = useState('');
  const [passwordLogin, setPasswordLogin] = useState('');

  const [isSignupShow, setIsSignupShow] = useState(false);
  const [isLoginShow, setIsLoginShow] = useState(false);

  if(props.token){
      axios
        .post('/token/validate', {
          token: props.token,
        })
        .then(async (response) => {
          router.push('/home');
        })
        .catch(err => {
        });
  }

  const createUser = (e) => {
    e.preventDefault();
    axios
      .post('/user/create', {
        name: nameSignUp,
        username: emailSignUp,
        password: passwordSignUp,
      })
      .then(async (response) => {
        setNameSignUp('');
        setEmailSignUp('');
        setPasswordSignUp('');
        toast.success('User created successfully, please login.', {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
      })
      .catch(err => {
        toast.error('User Already exists', err, {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
        console.log(err.toString());
      });
  };

  const validateUser = (e) => {
    e.preventDefault();
    axios
      .post('/user/login', {
        username: emailLogin,
        password: passwordLogin,
      })
      .then(async (response) => {
        cookie.set('token', response.data.token)
        cookie.set('user', JSON.stringify(response.data))
        setEmailLogin('');
        setPasswordLogin('');
        router.push('/home');
      })
      .catch((err) => {
        toast.error('Error logging in, ', err, {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
        console.error(err);
      });
  };

    return(
      <div className={styles.container}>
        <h4 className={styles.title}>SIGN UP!</h4>
        <form className={styles.main} onSubmit={(e)=>createUser(e)}>
          <input className={styles.text} type="text" placeholder="Full Name"
            value={nameSignUp} required
            onChange={(e)=>setNameSignUp(e.target.value)}
            />
           <input className={styles.text} type="email" placeholder="Email"
            autoComplete="email"
            value={emailSignUp} required
            onChange={(e)=>setEmailSignUp(e.target.value)}
            />
            <div className={styles.password}>
              <input className={styles.text} type={isSignupShow ? "text" : "password" } placeholder="Password"
              value={passwordSignUp} autoComplete="password" required
              onChange={(e)=>setPasswordSignUp(e.target.value)}
              /> 
              <span><FontAwesomeIcon className={styles.eye} size="1x" icon={isSignupShow?faEye:faEyeSlash} onClick={() => setIsSignupShow(!isSignupShow)}/></span>
            </div>
            <button className={styles.button} type="submit">Create
              <i className="material-icons right">forward</i>
            </button>
        </form>

        <h6 className={styles.text2}> OR </h6>
      
        <h4 className={styles.title}>LOGIN</h4>
        <form className={styles.main} onSubmit={(e)=>validateUser(e)}>
           <input className={styles.text} type="email" placeholder="Email"
            autoComplete="email" required
            value={emailLogin}
            onChange={(e)=>setEmailLogin(e.target.value)}
            />
            <div className={styles.password}>
              <input className={styles.text} type={isLoginShow ? "text" : "password" } placeholder="Password"
              value={passwordLogin} autoComplete="password" required
              onChange={(e)=>setPasswordLogin(e.target.value)}
              />
              <span><FontAwesomeIcon className={styles.eye} size="1x" icon={isLoginShow?faEye:faEyeSlash} onClick={() => setIsLoginShow(!isLoginShow)}/></span>
            </div>
            <button className={styles.button} type="submit">Login
              <i className="material-icons right">forward</i>
            </button>
        </form>
        <ToastContainer
        position="top-center"
        autoClose={3000}
        hideProgressBar
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
      />
      </div>
    )
}

export async function getServerSideProps(ctx){
  const {user, token} = parseCookies(ctx)
  return {props: {token: token || null, user: user || {}}};
}