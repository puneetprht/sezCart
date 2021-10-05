import cookie from 'js-cookie';
import {parseCookies} from 'nookies';
import axios from '../../src/service/axios';
import router, {useRouter} from 'next/router';
import 'react-toastify/dist/ReactToastify.css';
import {useState, useEffect, useRef} from 'react';
import { ToastContainer, toast } from 'react-toastify';

import styles from '../../styles/Home.module.css';

export default function Home(props) {
  if(props.token){
    axios
      .post('/token/validate', {
        token: props.token,
      })
      .then(async (response) => {
      })
      .catch(err => {
        cookie.remove('token')
        cookie.remove('user')
        router.push('/')
      });
  }

  const [itemList, setItemList] = useState([]);
  const [cartItemList, setCartItemList] = useState([]);
  const [user, setUser] = useState(JSON.parse(props.user));

  useEffect(() => {
    getItems();
    if(user && user.cart_id > 0){
      fetchCartList(user.cart_id)
    }
  }, []);

  const getItems = () => {
    axios
      .get('/item/list')
      .then((response) => {         
        if (response.data) {
          setItemList(response.data);
        }
      })
      .catch((err) => {
        console.error(err);
      });
  }

  function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
  }

  const postCartItem = (e, item) => {    
    e.preventDefault();
    let cartItem = {
      cart_id: user.cart_id,
      item_id: item.id
    }
    axios
      .post('/cart/add', cartItem, {
        headers: { Authorization: `Bearer ${props.token}` }
    })
      .then(async (response) => {
        console.log(response.data);
        if (user.cart_id == 0){
          user.cart_id = response.data.cart_id
          setUser(user);
        }
        fetchCartList(response.data.cart_id);
        toast.success('Item added to cart.', {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
      })
      .catch(err => {
        toast.error('There was some error adding item to the Cart.', {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
      });
  }

  const fetchCartList = (cartId) => {
    console.log("cartId: ",cartId);
    axios
      .get('/cart/list?cartId=' + cartId)
      .then((response) => {         
        if (response.data) {
          setCartItemList(response.data);
        }
      })
      .catch((err) => {
        console.error(err);
      });
  }

  const showItemName = (item_id) => {
    return itemList.filter(item => item.id == item_id).map(item => item.name)[0]
  }

  const showCartList = () => {
    if (user && user.cart_id && cartItemList.length) {
      toast.success(cartItemList.map(item => showItemName(item.item_id) + ' x ' + item.count).toString(), {
        position: "top-center",
        autoClose: 3000,
        hideProgressBar: true,
        closeOnClick: true,
        draggable: true,
      });
    } else{
      toast.error('Nothing to show.', {
        position: "top-center",
        autoClose: 3000,
        hideProgressBar: true,
        closeOnClick: true,
        draggable: true,
      });
    }
  }

  const showOrders = () => {
    axios
      .get('/order/list?userId=' + user.id)
      .then((response) => {         
        if (response.data.length > 0) {
          console.log(response.data);
          toast.success('OrderIds: ' + response.data.map(item => item.id).toString().toString(), {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        }); 
        } else {
          toast.error('Nothing to show.', {
            position: "top-center",
            autoClose: 3000,
            hideProgressBar: true,
            closeOnClick: true,
            draggable: true,
          });
        }
      })
      .catch((err) => {
        toast.error('Nothing to show.', {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
      });
  }

  const checkoutCart = (e) => {
    if(user.cart_id == 0 || cartItemList.length == 0) {
      toast.error('No Items to checkout.', {
        position: "top-center",
        autoClose: 3000,
        hideProgressBar: true,
        closeOnClick: true,
        draggable: true,
      });
    } else {
      axios
      .post('/cart/'+user.cart_id+'/complete', {}, {
        headers: { Authorization: `Bearer ${props.token}` }
      })
      .then(async (response) => {
        user.cart_id = 0
        setUser(user);
        setCartItemList([]);
        toast.success('Cart checked out.', {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
      })
      .catch(err => {
        toast.error('There was some error checking out Cart.', {
          position: "top-center",
          autoClose: 3000,
          hideProgressBar: true,
          closeOnClick: true,
          draggable: true,
        });
      });
    }
  }

  return (
    <div>
      <div className = {styles.buttonWrapper}>
                <span></span>
                <div>
                  <button className={styles.button2} type="submit" onClick={(e)=> checkoutCart(e)}>Checkout
                  </button>
                  <button className={styles.button2} type="submit" onClick={(e)=> showCartList(e)}>Cart Items
                  </button>
                  <button className={styles.button2} type="submit" onClick={(e) => showOrders(e)}>Orders
                  </button>
                  </div>
      </div>
      <main className={styles.container}>
        <div className={styles.centerModule}>
          <p className={styles.rightHeading}>All Items in the store</p>
          {
            itemList.map((item, index) => {
              return (
                <div key={item.id} className={styles.centerGrid}>
                  <p className={styles.centerGridSubject}>{index+1}. {capitalizeFirstLetter(item.name)}</p>
                  <button className={styles.button} type="submit" onClick={(e) => postCartItem(e, item)}>
                    Add to Cart
                  </button>
                </div>
            )})
          }  
        </div>
        <div className={styles.rightModule}>
          <p className={styles.rightHeading}>Cart</p>
          <p className={styles.rightContent}>
          {
            cartItemList.map((item, index) => {
              return (
                <div key={item.id}>
                  <p>{index+1}. {showItemName(item.item_id)}  x   {item.count}</p>
                </div>
            )})
          }
          </p>
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

export async function getServerSideProps(ctx){
  const {user, token} = parseCookies(ctx)
  if(!token){
      const {res} = ctx
      res.writeHead(302,{Location:"/"})
      res.end()
  }
  return {props: {token: token, user: user || {}}};
}