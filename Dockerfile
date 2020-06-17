FROM scratch
LABEL authors="Parikshit Gothwal <parikshit@urantiatech.com>"
ADD parikshitg parikshitg
ADD static static
ADD views views
EXPOSE 8080
CMD ["/parikshitg"]
