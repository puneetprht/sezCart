import _ from 'lodash';
import Head from 'next/head';
import Link from 'next/link';
import {parseCookies} from 'nookies';
import {useRouter} from 'next/router';
import axios from '../../src/service/axios';
import 'react-toastify/dist/ReactToastify.css';
import {useState, useEffect, useRef} from 'react';
import { ToastContainer, toast } from 'react-toastify';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheck, faTimes, faEdit, faPlus, faTrashAlt } from '@fortawesome/free-solid-svg-icons';

import styles from '../../styles/Home.module.css';

export default function Home(props) {
  const router = useRouter();

  const [user, setUser] = useState(JSON.parse(props.user));

  const [centralItem, setCentralItem] = useState('feed');

  const displayToasterMessage = (toastType, message) => {
    if(toastType = 'error'){
      toast.error(Message, {
        position: "top-center",
        autoClose: 3000,
        hideProgressBar: true,
        closeOnClick: true,
        draggable: true,
      });
    } else if (toastType = 'success') {
        toast.success(Message, {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
    } else if (toastType = 'warning') {
        toast.warning(Message, {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
    } else if (toastType = 'info') {
      toast.info(Message, {
        position: "top-center",
        autoClose: 3000,
        hideProgressBar: true,
        closeOnClick: true,
        draggable: true,
      });
    } else if (toastType = 'dark') {
      toast.dark(Message, {
        position: "top-center",
        autoClose: 3000,
        hideProgressBar: true,
        closeOnClick: true,
        draggable: true,
      });
    } else {
        toast.default(Message, {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
    }
  }

  return (
    <div className={styles.wrapper}>
      <main className={styles.container}>
        <div className={styles.centerModule}>
          List of all Items with an add button  
        </div>
        <div className={styles.rightModule}>
          <p className={styles.rightHeading}>Cart</p>
          <p className={styles.rightContent}>Update all items here.</p>
        </div>
      </main>
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