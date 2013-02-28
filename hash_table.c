#define HASHSIZE 256


// Hash table.
// It's an array with index as hash, 
// value as a pointer to an array of entries.
void* hash_table[];

void* lookup(char *s){
    void *p;
    if ((*p = hash_table[hash(*s)]) != NULL) {
        if(strcmp(*p, ))
    }
    for (
        p = hash_table[hash(*s)];
        p != NULL;
        
        )
    return NULL;
}

// Return a number from *s. 
unsigned hash(char *s){
    unsigned result;
    for(result = 0; *s != '\0'; s++) {
        result = *s + result * 31; 
    }
    return result % HASHSIZE;
}



