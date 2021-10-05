import Link from 'next/link';
import cookie from 'js-cookie';
import {parseCookies} from 'nookies';
import {useRouter} from 'next/router';

import styles from '../../styles/Navbar.module.css';

const NavBar = () => {
    const router = useRouter();
    const cookieuser = parseCookies();
    const token =  cookieuser.token;
    
    return(
      <>
        <nav className={styles.navContainer}>
          <div className={styles.wrapper}>
          <div className={styles.logo}>
              <Link  href="/">
                SezCart
              </Link>
          </div>
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