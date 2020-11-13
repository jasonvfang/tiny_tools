#include <stdio.h>
#include <sys/socket.h>
#include <net/if.h>
#include <arpa/inet.h>
#include <netinet/in.h>
#include <string.h>
#include <stdlib.h>

static inline int addrconf_ifid_eui48(uint8_t *eui, uint8_t *dev_mac_addr)
{

	memcpy(eui, dev_mac_addr, 3);
	memcpy(eui + 5, dev_mac_addr + 3, 3);

	eui[3] = 0xFF;
	eui[4] = 0xFE;
	eui[0] ^= 2;

	return 0;
}

int generate_ipv6_link_local_address(uint8_t dev_mac[6], struct in6_addr *out_ipv6_link_local_addr)
{
    struct in6_addr addr;

    addr.s6_addr32[0] = htonl(0xfe800000);
    addr.s6_addr32[1] = 0;

    addrconf_ifid_eui48(addr.s6_addr + 8, dev_mac);

    if(out_ipv6_link_local_addr)
    {
        memcpy(out_ipv6_link_local_addr, &addr, sizeof(struct in6_addr));
    }
    
    return 0;
}

int test_ipv6_exist(void)
{
    int exist = 0;
    FILE *fp = NULL;
    fp = popen("ifconfig eth0 | grep inet6", "r");
    if(fp)
    {
        char buffer[128];
        memset(buffer, 0, sizeof(buffer));
        fgets(buffer, sizeof(buffer) - 1, fp);

        if (strlen(buffer) > 0 && buffer[strlen(buffer) - 1] == '\n') {
            buffer[strlen(buffer) - 1] = '\0';
        }
        
        if(strlen(buffer) >0 && strstr(buffer, "fe80::") != NULL){
            exist = 1;
        }          

        printf("len:%d, %s\n", strlen(buffer), buffer);

        pclose(fp);

    }
    
    return exist;
}

int main(void)
{
    struct in6_addr ipv6_addr;
    int i = 0;

    //00:0E:C6:DF:8C:AA
    uint8_t mac[6] = {0x00, 0x0E, 0xC6, 0xDF, 0x8C, 0xAA};
    
    printf("MAC address: 0x%02X-0x%02X-0x%02X-0x%02X-0x%02X-0x%02X\n", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5]);

    generate_ipv6_link_local_address(mac, &ipv6_addr);

    //dump values
    printf("IPv6 Link Local Address:\n");
    for(; i < 16; i ++)
    {
        printf("0x%02X ", ipv6_addr.s6_addr32[i]);
    }
    printf("\n");

    if(test_ipv6_exist()){
        printf("ipv6 addr for eth0 is configured\n");
        return 0;
    }

    char ipv6_str_buf[INET6_ADDRSTRLEN] = {0 };
    inet_ntop(AF_INET6, &(ipv6_addr), ipv6_str_buf, sizeof(ipv6_str_buf));

    printf("String:%s\n", ipv6_str_buf);

    //Set ipv6 address for eth0
    char cmd[128] = {0};
    snprintf(cmd, sizeof(cmd) - 1, "ip -6 addr add %s/64 dev %s", ipv6_str_buf, "eth0");
    system(cmd);

    return 0;
}


