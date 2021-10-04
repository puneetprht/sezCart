import Link from 'next/link';
import cookie from 'js-cookie';
import {useState} from 'react';
import {parseCookies} from 'nookies';
import {useRouter} from 'next/router';

import styles from '../../styles/Navbar.module.css';

const NavBar = () => {
    const router = useRouter();
    const cookieuser = parseCookies();
    const token =  cookieuser.token;
    const [mode, setMode] = useState(false);


    const overlayMode = () => {
      if(mode)
        return styles.overlayActive;
      else
        return styles.overlay;
    } 
    
    return(
      <>
        <nav className={styles.navContainer}>
          <div className={styles.wrapper}>
          {token ?
              <>
                <a id="button-laptop">
                  <button className={styles.logout} onClick={()=>{
                      cookie.remove('token')
                      cookie.remove('user')
                      router.push('/')
                    }}
                  >
                    Log Out
                  </button>
                </a>   
              </>
                :
              <>
              </>
            }
            <Link href="/">
                SezCart
            </Link>
            {token ?
              <>
                <a id="button-laptop">
                  <button className={styles.logout} onClick={()=>{
                      cookie.remove('token')
                      cookie.remove('user')
                      router.push('/')
                    }}
                  >
                    Log Out
                  </button>
                </a>   
              </>
                :
              <>
              </>
            }
          </div>
        </nav>
      </>
    )
}

export default NavBar