import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import Modal from "../components/modal";

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>SlackClone</title>
        <meta name="description" content="slack clone app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <div className={styles.groupMenu}>aaaaa</div>
        <div className={styles.chatScreen}>aaaab</div>
      </main>
    </div>
  );
}
