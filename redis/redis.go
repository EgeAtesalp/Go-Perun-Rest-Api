package redis

import (

	"context"

	"restapidemo/model"
	"strconv"
	"strings"
	"crypto/rand"
	"io"
	"golang.org/x/crypto/nacl/secretbox"
	"encoding/hex"

	"github.com/go-redis/redis/v8"
	

	


	
)

//Creates new database client, can be customized later 
//to add privacy options or other database options
func CreateNewDatabase(addr string)(*redis.Client){
	
	rdb := redis.NewClient(&redis.Options{
		Addr:	  addr,
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
	return rdb
}

//Takes model.Usersession as argument,adds new user to the database, 
//overrides if another session with the same alias already exists
func AddOrUpdateUserSession(rdb *redis.Client,usersess model.UserSession)(error){
	ctx := context.Background()
	val := ConvertToString(usersess)
	err := rdb.Set(ctx, usersess.Alias, val , 0).Err()
	return err
}


//Takes alias as argument, finds the session with the alias as the key,
//returns error and empty user session if not found
func FindUserSession(rdb *redis.Client, alias string, )(model.UserSession, error){
	ctx := context.Background()
	val, err := rdb.Get(ctx, alias).Result()
	if err != nil {
		
		nilsess := model.UserSession{
			Alias:   "",
			Ip_Addr: "",
			Port:    0,
		} 
		return nilsess, err
	}
	response := ConvertToUserSession(val,alias)

	return response, nil
}

//Takes alias as argument, deletes the session from database
func DeleteUserSession(rdb *redis.Client,alias string)(error){
	ctx := context.Background()
	
	err := rdb.Del(ctx, alias ).Err()
	return err
}

//Helper function, takes stored string value from database and converts it
//to model.UserSession
func ConvertToUserSession(val string, alias string)(model.UserSession){

	s := strings.Split(val, "/")
	var ip_address = s[0]
	var portstr = s[1]
	i, err := strconv.ParseInt(portstr, 10, 16)
	if err != nil {
   		 panic(err)
	}
	port := int16(i)
	var response = model.UserSession{Alias: alias, Ip_Addr: ip_address, Port: port}
	return response
}

//Helper function, takes model.UserSession and converts it
//into string to store in database
func ConvertToString(sess model.UserSession)(string){

	var ip_address = sess.Ip_Addr
	var port = sess.Port
	portstr := strconv.FormatInt(int64(port), 10)
	
	response := ip_address + "/" + portstr
	return response
}


//Encryption Library

//Takes model.Usersession as argument,adds encrypted new user to the database, 
//overrides if another session with the same alias already exists
func AddOrUpdateEncryptedUserSession(rdb *redis.Client,secretKey [32]byte,usersess model.UserSession)(error){
	ctx := context.Background()
	val := EncryptAndConvertToByteArray(usersess,secretKey)
	err := rdb.Set(ctx, usersess.Alias, val , 0).Err()
	return err
}

//Takes alias as argument, finds the session with the alias as the key,
//returns decrypted user session if succesful,
//returns error and empty user session if not found
func FindEncryptedUserSession(rdb *redis.Client, secretKey [32]byte,alias string, )(model.UserSession, error){
	ctx := context.Background()
	val, err := rdb.Get(ctx, alias).Bytes()
	if err != nil {
		
		nilsess := model.UserSession{
			Alias:   "",
			Ip_Addr: "",
			Port:    0,
		} 
		return nilsess, err
	}
	response := DecryptAndConvertToUserSession(val,alias,secretKey)

	return response, nil
}

//Helper function, takes byte array, 
//decrypts and converts into model.UserSession
func DecryptAndConvertToUserSession(val []byte, alias string, secretKey [32]byte)(model.UserSession){
	
	decryptedByteArray := string(Decrypt(secretKey,val))



	s := strings.Split(decryptedByteArray, "/")
	var ip_address = s[0]
	var portstr = s[1]
	i, err := strconv.ParseInt(portstr, 10, 16)
	if err != nil {
   		 panic(err)
	}
	port := int16(i)
	var response = model.UserSession{Alias: alias, Ip_Addr: ip_address, Port: port}
	return response
}

//helper function, takes model.UserSession
//encrypts and converts into byte array
func EncryptAndConvertToByteArray(sess model.UserSession, secretKey [32]byte)([]byte){

	var ip_address = sess.Ip_Addr
	var port = sess.Port
	portstr := strconv.FormatInt(int64(port), 10)
	response := []byte(ip_address + "/" + portstr)

	encryptedResponse := Encrypt(secretKey,response)
	return encryptedResponse
}

//encryption function using SecretBox
//takes byte array and key,
//nonce is stored in the message itself
func Encrypt(key [32]byte, message []byte)([]byte){
   	
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}

	encrypted := secretbox.Seal(nonce[:], message, &nonce, &key)

	return encrypted
}

//decryption function using SecretBox,
//takes bute array and key,
//nonce is extracted from the byte array
func Decrypt(key [32]byte, message []byte)([]byte){
   	

	var decryptNonce [24]byte
	copy(decryptNonce[:], message[:24])
	decrypted, ok := secretbox.Open(nil, message[24:], &decryptNonce, &key)
	if !ok {
		panic("decryption error")
	}
	return decrypted
}

//creates secret key using a string
//String must be of even length, and a stringfied version of a hexadecimal
func CreateSecretKey(s string)([32]byte){
	secretKeyBytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	var secretKey [32]byte
	copy(secretKey[:], secretKeyBytes)
	return secretKey
}