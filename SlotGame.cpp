#include <iostream>
#include <cstdlib>
#include <cmath>
#include <stdbool.h>
#include <vector>
#include <thread>
#include <chrono>
#include<iomanip>
#include <string.h>
#include <unistd.h>

using namespace std;

enum
{
    SLOT_TYPE_A = 0,
    SLOT_TYPE_B,
    SLOT_TYPE_C,
};

enum
{
    LOSE = 0,
    WIN = 1,
};

#define SLEEP_TIME_SLOT_A 50  //milliseconds
#define SLEEP_TIME_SLOT_B 150  //milliseconds
#define SLEEP_TIME_SLOT_C 250  //milliseconds

const std::vector < std::vector < std::string > > asciiNumbers =
{
    {" ___ "," / _ \\  "," | | | | "," | | | | ", " | |_| | "," \\___/  ",""}, // 0
    {" __ "," /_ | ", " | | "," | | ", " | | "," |_| ",""}, // 1
    {" ___  "," |__ \\ "," ) | "," / / ", " / /_ "," |____| ",""}, // 2
    {" ____ "," |___ \\ "," __) | "," |__ < ", " ___) | "," |____/ ",""}, // 3
    {" _   _  "," | |  | |  ", " | |  | |_"," |__ _|   ", "    | |  ","    |_|  ",""}, // 4
    {" _____ "," | ____| ", " | |__   "," |___ \\ "," ___) | "," |____/ ",""}, // 5
    {" __   ","/ /    ", "/ /_    "," | '_ \\  ","  | (_) | "," \\___/  ",""}, // 6
    {" ______ "," |____ |", " / / "," / / ", " / / "," /_/ ",""}, // 7
    {" ___   "," / _ \\  "," | (_) | "," > _ < ", " | (_) | "," \\___/  ",""}, // 8
    {" ___ "," / _ \\ "," | (_) | "," \\__, | ", " / / "," /_/ ",""}, // 9
};


class SlotGame
{
private:
    /* data */

public:
    int SlotNumberA;
    int SlotNumberB;
    int SlotNumberC;
    int total_win_bets;

    int playTheGame();

    void DigitsSpin(int num);

    void genRandomSlot(int SlotType);

    int loseOrWin();

    void mySleepInMilliseconds(unsigned int milliseconds)
    {
         std::this_thread::sleep_for(std::chrono::milliseconds(milliseconds));
    }

    SlotGame();

    ~SlotGame();
};


void SlotGame::DigitsSpin(int num)
{
    //cout << "num == > " << num << endl;
    int A = num / 100;
    int B = (num / 10) % 10;
    int C = num % 10;

    char L1[64], L2[64], L3[64], L4[64], L5[64], L6[64], L7[64];

    sprintf(L1, "%11s%11s%11s", asciiNumbers[A][0].c_str(), asciiNumbers[B][0].c_str(), asciiNumbers[C][0].c_str());
    sprintf(L2, "%11s%11s%11s", asciiNumbers[A][1].c_str(), asciiNumbers[B][1].c_str(), asciiNumbers[C][1].c_str());
    sprintf(L3, "%11s%11s%11s", asciiNumbers[A][2].c_str(), asciiNumbers[B][2].c_str(), asciiNumbers[C][2].c_str());
    sprintf(L4, "%11s%11s%11s", asciiNumbers[A][3].c_str(), asciiNumbers[B][3].c_str(), asciiNumbers[C][3].c_str());
    sprintf(L5, "%11s%11s%11s", asciiNumbers[A][4].c_str(), asciiNumbers[B][4].c_str(), asciiNumbers[C][4].c_str());
    sprintf(L6, "%11s%11s%11s", asciiNumbers[A][5].c_str(), asciiNumbers[B][5].c_str(), asciiNumbers[C][5].c_str());
    sprintf(L7, "%11s%11s%11s", asciiNumbers[A][6].c_str(), asciiNumbers[B][6].c_str(), asciiNumbers[C][6].c_str());

    cout << L1 << endl << L2 << endl << L3 << endl << L4 << endl << L5 << endl << L6 << endl << L7 << endl;
}


void SlotGame::genRandomSlot(int SlotType)
{
    int num = 0;
    time_t t1, t2;

    if (SlotType == SLOT_TYPE_A)
    {
        //cout << "Gen Slot A" << endl;
        time(&t1);

        while (1)
        {
            time(&t2);

            if (t2 - t1 >= 8)
                break;

            printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n");

            num = rand() % 899 + 100;
            DigitsSpin(num);

            printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n");

            mySleepInMilliseconds(SLEEP_TIME_SLOT_A);

            system("clear");            

        }

        SlotNumberA = num / 100;
    }
    else if (SlotType == SLOT_TYPE_B)
    {
        //cout << "Gen Slot B" << endl;
        time(&t1);

        while (1)
        {
            time(&t2);

            if (t2 - t1 >= 5)
                break;          

            printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n");

            num = rand() % 99 + SlotNumberA * 100;
            DigitsSpin(num);

            printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n");
            
            mySleepInMilliseconds(SLEEP_TIME_SLOT_B);
            system("clear");     
        }

        SlotNumberB = (num / 10) % 10;
    }
    else if (SlotType == SLOT_TYPE_C)
    {
        //cout << "Gen Slot C" << endl;
        time(&t1);

        while (1)
        {
            printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n");
            
            num = rand() % 9 + SlotNumberA * 100 + SlotNumberB * 10;
            DigitsSpin(num);
            
            mySleepInMilliseconds(SLEEP_TIME_SLOT_C);

            time(&t2);

            if (t2 - t1 >= 3)
                break;
            
            printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n");
            system("clear");  
        }

        SlotNumberC = num % 10;
    }
}


int SlotGame::loseOrWin()
{
    int res = LOSE;

    //cout << "slot is " << SlotNumberA << SlotNumberB << SlotNumberC <<endl;

    if (SlotNumberA == SlotNumberB == SlotNumberC)
    {
        res = WIN;
        total_win_bets += 10;
        cout<<"You win, total " << total_win_bets << "\n\n";
    }
    else
    {
        if(SlotNumberA == SlotNumberB
            || SlotNumberB == SlotNumberC
            || SlotNumberA == SlotNumberC)
        {
            res = WIN;
            total_win_bets += 5;
            cout<<"You win 5 more bets, total " << total_win_bets << "\n\n";
        }
        else
        {
            res = LOSE;
            total_win_bets -= 1;
            cout<<"You lose, total bets " << total_win_bets << "\n\n";
        }
    }

    return res;
}


int SlotGame::playTheGame()
{
    int bet = 0;

    while (1)
    {
        cout<<"Enter your bet as a whole numbers ? ";
        cin>>bet;

        if (bet < 0 && bet >= 1000)
        {
            cout << "Invalid bet numbers, Must be 001 - 999" << endl;
            continue;
        }

        genRandomSlot(SLOT_TYPE_A);
        genRandomSlot(SLOT_TYPE_B);
        genRandomSlot(SLOT_TYPE_C);

        loseOrWin();

        if (total_win_bets <= 0)
        {
            cout << "You lost $" << bet << endl;
            break;
        }
    }

    return 0;
}


SlotGame::SlotGame()
{
    SlotNumberA = 0;
    SlotNumberB = 0;
    SlotNumberC = 0;
    total_win_bets = 0;
}


SlotGame::~SlotGame()
{

}

int main(int argc, char **argv)
{
    SlotGame *A = new SlotGame;

    A->playTheGame();

    delete A;

    return 0;
}
