FROM scratch
LABEL authors="Sangeet Kumar <sk@urantiatech.com>"
ADD parikshitg parikshitg
ADD static static
ADD views views
EXPOSE 8080
CMD ["/parikshitg"]
